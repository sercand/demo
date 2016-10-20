// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: color.proto

#ifndef PROTOBUF_color_2eproto__INCLUDED
#define PROTOBUF_color_2eproto__INCLUDED

#include <string>

#include <google/protobuf/stubs/common.h>

#if GOOGLE_PROTOBUF_VERSION < 3000000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please update
#error your headers.
#endif
#if 3000000 < GOOGLE_PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)

namespace feed {
namespace v1 {

// Internal implementation detail -- do not call these.
void protobuf_AddDesc_color_2eproto();
void protobuf_AssignDesc_color_2eproto();
void protobuf_ShutdownFile_color_2eproto();

class Color;
class NewColorRequest;

// ===================================================================

class Color : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:feed.v1.Color) */ {
 public:
  Color();
  virtual ~Color();

  Color(const Color& from);

  inline Color& operator=(const Color& from) {
    CopyFrom(from);
    return *this;
  }

  static const ::google::protobuf::Descriptor* descriptor();
  static const Color& default_instance();

  void Swap(Color* other);

  // implements Message ----------------------------------------------

  inline Color* New() const { return New(NULL); }

  Color* New(::google::protobuf::Arena* arena) const;
  void CopyFrom(const ::google::protobuf::Message& from);
  void MergeFrom(const ::google::protobuf::Message& from);
  void CopyFrom(const Color& from);
  void MergeFrom(const Color& from);
  void Clear();
  bool IsInitialized() const;

  int ByteSize() const;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input);
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* output) const;
  ::google::protobuf::uint8* SerializeWithCachedSizesToArray(::google::protobuf::uint8* output) const {
    return InternalSerializeWithCachedSizesToArray(false, output);
  }
  int GetCachedSize() const { return _cached_size_; }
  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const;
  void InternalSwap(Color* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return _internal_metadata_.arena();
  }
  inline void* MaybeArenaPtr() const {
    return _internal_metadata_.raw_arena_ptr();
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // optional string hex = 1;
  void clear_hex();
  static const int kHexFieldNumber = 1;
  const ::std::string& hex() const;
  void set_hex(const ::std::string& value);
  void set_hex(const char* value);
  void set_hex(const char* value, size_t size);
  ::std::string* mutable_hex();
  ::std::string* release_hex();
  void set_allocated_hex(::std::string* hex);

  // @@protoc_insertion_point(class_scope:feed.v1.Color)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  bool _is_default_instance_;
  ::google::protobuf::internal::ArenaStringPtr hex_;
  mutable int _cached_size_;
  friend void  protobuf_AddDesc_color_2eproto();
  friend void protobuf_AssignDesc_color_2eproto();
  friend void protobuf_ShutdownFile_color_2eproto();

  void InitAsDefaultInstance();
  static Color* default_instance_;
};
// -------------------------------------------------------------------

class NewColorRequest : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:feed.v1.NewColorRequest) */ {
 public:
  NewColorRequest();
  virtual ~NewColorRequest();

  NewColorRequest(const NewColorRequest& from);

  inline NewColorRequest& operator=(const NewColorRequest& from) {
    CopyFrom(from);
    return *this;
  }

  static const ::google::protobuf::Descriptor* descriptor();
  static const NewColorRequest& default_instance();

  void Swap(NewColorRequest* other);

  // implements Message ----------------------------------------------

  inline NewColorRequest* New() const { return New(NULL); }

  NewColorRequest* New(::google::protobuf::Arena* arena) const;
  void CopyFrom(const ::google::protobuf::Message& from);
  void MergeFrom(const ::google::protobuf::Message& from);
  void CopyFrom(const NewColorRequest& from);
  void MergeFrom(const NewColorRequest& from);
  void Clear();
  bool IsInitialized() const;

  int ByteSize() const;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input);
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* output) const;
  ::google::protobuf::uint8* SerializeWithCachedSizesToArray(::google::protobuf::uint8* output) const {
    return InternalSerializeWithCachedSizesToArray(false, output);
  }
  int GetCachedSize() const { return _cached_size_; }
  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const;
  void InternalSwap(NewColorRequest* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return _internal_metadata_.arena();
  }
  inline void* MaybeArenaPtr() const {
    return _internal_metadata_.raw_arena_ptr();
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // @@protoc_insertion_point(class_scope:feed.v1.NewColorRequest)
 private:

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  bool _is_default_instance_;
  mutable int _cached_size_;
  friend void  protobuf_AddDesc_color_2eproto();
  friend void protobuf_AssignDesc_color_2eproto();
  friend void protobuf_ShutdownFile_color_2eproto();

  void InitAsDefaultInstance();
  static NewColorRequest* default_instance_;
};
// ===================================================================


// ===================================================================

#if !PROTOBUF_INLINE_NOT_IN_HEADERS
// Color

// optional string hex = 1;
inline void Color::clear_hex() {
  hex_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& Color::hex() const {
  // @@protoc_insertion_point(field_get:feed.v1.Color.hex)
  return hex_.GetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void Color::set_hex(const ::std::string& value) {
  
  hex_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:feed.v1.Color.hex)
}
inline void Color::set_hex(const char* value) {
  
  hex_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:feed.v1.Color.hex)
}
inline void Color::set_hex(const char* value, size_t size) {
  
  hex_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:feed.v1.Color.hex)
}
inline ::std::string* Color::mutable_hex() {
  
  // @@protoc_insertion_point(field_mutable:feed.v1.Color.hex)
  return hex_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* Color::release_hex() {
  // @@protoc_insertion_point(field_release:feed.v1.Color.hex)
  
  return hex_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void Color::set_allocated_hex(::std::string* hex) {
  if (hex != NULL) {
    
  } else {
    
  }
  hex_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), hex);
  // @@protoc_insertion_point(field_set_allocated:feed.v1.Color.hex)
}

// -------------------------------------------------------------------

// NewColorRequest

#endif  // !PROTOBUF_INLINE_NOT_IN_HEADERS
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace v1
}  // namespace feed

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_color_2eproto__INCLUDED
